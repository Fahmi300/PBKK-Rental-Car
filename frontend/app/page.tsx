"use client";

import React, { useEffect, useState } from "react";
import Navbar from "../components/navbar";
import { useRouter } from "next/navigation";

interface Car {
  id: number;
  name: string;
  brand: string;
  year: number;
  price_per_day: number;
  availability: boolean;
  imageUrl?: string; // Converted URL for rendering
  category?: {
    id: number;
    name: string;
  };
}

const DashboardPage = () => {
  const [cars, setCars] = useState<Car[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  useEffect(() => {
    const fetchCars = async () => {
      try {
        setLoading(true);
        const response = await fetch("http://localhost:8080/api/car");
        if (!response.ok) {
          throw new Error("Failed to fetch cars");
        }
        const data = await response.json();

        // Process cars and fetch image blobs
        const carsWithImages = await Promise.all(
          data.data.map(async (car: Car) => {
            const imageUrl = `http://localhost:8080/api/car/${car.id}/image`; // Set the URL for the image
            return { ...car, imageUrl }; // Add `imageUrl` for rendering
          })
        );

        setCars(carsWithImages);
      } catch (err: any) {
        setError(err.message || "An error occurred");
      } finally {
        setLoading(false);
      }
    };

    fetchCars();
  }, []);

  useEffect(() => {
    // Cleanup: Revoke object URLs when component unmounts or when `cars` state changes
    return () => {
      cars.forEach((car) => {
        if (car.imageUrl) {
          URL.revokeObjectURL(car.imageUrl);
        }
      });
    };
  }, [cars]);

  if (loading) {
    return <p className="text-center text-gray-600 mt-10">Loading cars...</p>;
  }

  if (error) {
    return (
      <p className="text-center text-red-600 mt-10">
        Error loading cars: {error}
      </p>
    );
  }

  console.log(cars)
  const handleRentNow = (carID: number) => {
    // Redirect to booking page and pass the carID in the URL
    router.push(`/booking?carID=${carID}`);
  };

  return (
    <div>
      <Navbar />
      <main className="p-6 bg-gray-100 min-h-screen">
        <div className="mb-6">
          <h1 className="text-3xl font-bold text-gray-800">Available Cars for Rent</h1>
          <p className="text-gray-600 mt-2">
            Browse through our collection of rentable cars. Check availability and book now!
          </p>
        </div>
        <section className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          {cars.map((car) => (
            <div key={car.id} className="bg-white shadow-md rounded-lg overflow-hidden">
              {car.imageUrl ? (
                <img src={car.imageUrl} alt={car.name} className="w-full h-48 object-cover" />
              ) : (
                <div className="w-full h-48 bg-gray-300 flex items-center justify-center">
                  <span className="text-gray-600">No Image Available</span>
                </div>
              )}
              <div className="p-4">
                <h2 className="text-xl font-bold text-gray-800">{car.name}</h2>
                <p className="text-gray-600 text-sm mt-1">
                  <strong>Brand:</strong> {car.brand}
                </p>
                <p className="text-gray-600 text-sm mt-1">
                  <strong>Year:</strong> {car.year}
                </p>
                <p className="text-gray-600 text-sm mt-1">
                  <strong>Category:</strong> {car.category?.name || "Uncategorized"}
                </p>
                <p className={`text-sm font-semibold mt-2 ${car.availability ? "text-green-600" : "text-red-600"}`}>
                  {car.availability ? "Available" : "Not Available"}
                </p>
                <p className="text-gray-800 text-lg font-semibold mt-4">
                  ${car.price_per_day.toFixed(2)} / day
                </p>
                <button
                  className={`mt-4 w-full py-2 rounded shadow-md ${car.availability ? "bg-blue-600 text-white hover:bg-blue-700" : "bg-gray-400 text-white cursor-not-allowed"}`}
                  onClick={() => car.availability && handleRentNow(car.id)} // Navigate to booking page when clicked
                  disabled={!car.availability}
                >
                  {car.availability ? "Rent Now" : "Unavailable"}
                </button>
              </div>
            </div>
          ))}
        </section>
      </main>
    </div>
  );
};

export default DashboardPage;
