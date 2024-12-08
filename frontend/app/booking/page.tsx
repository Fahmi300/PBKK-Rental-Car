'use client';

import React, { useState } from "react";
import Navbar from "../../components/navbar"; // Import Navbar component
import { useRouter, useSearchParams } from "next/navigation";

const BookingPage = () => {
  const searchParams = useSearchParams();
  const carID = searchParams?.get("carID"); // Extract carID from the URL
  const router = useRouter();

  const [formData, setFormData] = useState({
    startDate: "",
    endDate: "",
    description: "",
    location: "",
    need: "",
    phone: "",
  });

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [totalPrice, setTotalPrice] = useState<number | null>(null); // Store the calculated total price

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!carID) {
      alert("Car ID is missing. Please select a car.");
      return;
    }

    try {
      setLoading(true);
      const response = await fetch("http://localhost:8080/api/booking", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("token")}`, // Adjust based on your auth setup
        },
        body: JSON.stringify({
            car_id: parseInt(carID), 
            start_date: new Date(formData.startDate).toISOString(),
            end_date: new Date(formData.endDate).toISOString(),
            description: formData.description,
            location: formData.location,
            need: formData.need,
            phone: formData.phone,
          }),
      });

      if (!response.ok) {
        const error = await response.json();
        setError(error.message || "Failed to create booking");
        return;
      }

      const result = await response.json();
      setTotalPrice(result.total_price); // Update the total price received from the backend
      alert("Booking created successfully!");
      router.push(`/booking/${result.data.id}`); // Redirect to booking details page
    } catch (error) {
      console.error("Error creating booking:", error);
      setError("An unexpected error occurred. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <Navbar /> {/* Include the Navbar */}
      <main className="p-6 bg-gray-100 min-h-screen">
        {/* Header */}
        <div className="mb-6">
          <h1 className="text-3xl font-bold text-gray-800">Create a Booking</h1>
          <p className="text-gray-600 mt-2">
            Fill in the required details to book your car rental. Ensure the information is accurate!
          </p>
        </div>

        {/* Booking Form */}
        <form
          onSubmit={handleSubmit}
          className="bg-white p-6 rounded shadow-md max-w-xl mx-auto"
        >
          {error && (
            <p className="text-red-600 text-sm mb-4">
              Error: {error}
            </p>
          )}
          <label className="block mb-4">
            <span className="text-gray-700 font-semibold">Start Date:</span>
            <input
              type="date"
              name="startDate"
              value={formData.startDate}
              onChange={handleChange}
              className="block w-full mt-1 border border-gray-300 rounded-md shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
              required
            />
          </label>
          <label className="block mb-4">
            <span className="text-gray-700 font-semibold">End Date:</span>
            <input
              type="date"
              name="endDate"
              value={formData.endDate}
              onChange={handleChange}
              className="block w-full mt-1 border border-gray-300 rounded-md shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
              required
            />
          </label>
          <label className="block mb-4">
            <span className="text-gray-700 font-semibold">Description:</span>
            <textarea
              name="description"
              value={formData.description}
              onChange={handleChange}
              className="block w-full mt-1 border border-gray-300 rounded-md shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
              placeholder="Additional details or special requests..."
            />
          </label>
          <label className="block mb-4">
            <span className="text-gray-700 font-semibold">Location:</span>
            <input
              type="text"
              name="location"
              value={formData.location}
              onChange={handleChange}
              className="block w-full mt-1 border border-gray-300 rounded-md shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
              required
              placeholder="Pickup or delivery location"
            />
          </label>
          <label className="block mb-4">
            <span className="text-gray-700 font-semibold">Need:</span>
            <input
              type="text"
              name="need"
              value={formData.need}
              onChange={handleChange}
              className="block w-full mt-1 border border-gray-300 rounded-md shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
              placeholder="Specify your needs"
            />
          </label>
          <label className="block mb-4">
            <span className="text-gray-700 font-semibold">Phone:</span>
            <input
              type="tel"
              name="phone"
              value={formData.phone}
              onChange={handleChange}
              className="block w-full mt-1 border border-gray-300 rounded-md shadow-sm focus:border-blue-500 focus:ring focus:ring-blue-200"
              required
              placeholder="Your contact number"
            />
          </label>
          <button
            type="submit"
            className={`mt-4 w-full py-2 rounded shadow-md text-white ${
              loading ? "bg-gray-500 cursor-not-allowed" : "bg-blue-600 hover:bg-blue-700"
            }`}
            disabled={loading}
          >
            {loading ? "Processing..." : "Create Booking"}
          </button>

          {/* Display Total Price */}
          {totalPrice !== null && (
            <div className="mt-4 text-lg font-semibold">
              <p>Total Price: ${totalPrice}</p>
            </div>
          )}
        </form>
      </main>
    </div>
  );
};

export default BookingPage;
