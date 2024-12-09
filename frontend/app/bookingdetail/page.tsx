'use client';

import React, { useEffect, useState } from "react";
import { useAuthRedirect } from "@/lib/auth";
import Navbar from "@/components/navbar"; // Import Navbar component
import { useParams } from "react-router-dom";

interface BookingDetails {
  data: {
    id: number;
    startDate: string;
    endDate: string;
    totalPrice: number;
    description: string;
    location: string;
    need: string;
    phone: string;
    userID: number;
    carID: number;
  };
}

const BookingDetailPage: React.FC = () => {
  useAuthRedirect(); // Check authentication

  const [booking, setBooking] = useState<BookingDetails | null>(null);
  const { id } = useParams<{ id: string }>(); // Get booking ID from URL

  useEffect(() => {
    if (!id) {
      console.error("Booking ID is missing in the URL");
      return;
    }

    const fetchBookingDetails = async () => {
      try {
        const token = localStorage.getItem("token");
        const response = await fetch(`http://localhost:8080/api/booking/${id}`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          throw new Error("Failed to fetch booking details");
        }

        const data: BookingDetails = await response.json();
        setBooking(data);
      } catch (error) {
        console.error(error);
      }
    };

    fetchBookingDetails();
  }, [id]); // Dependency array ensures useEffect runs when the ID changes

  if (!booking) {
    return <div className="text-center mt-10 text-lg text-gray-700">Loading booking details...</div>; // Loader while data is not available
  }
  console.log(booking);

  return (
    <>
      <Navbar />
      <div className="flex flex-col items-center p-8 font-sans">
        <h1 className="text-3xl font-bold mb-6">Booking Details</h1>
        <div className="border border-gray-300 rounded-lg p-6 shadow-md w-full max-w-md bg-gray-50">
          <p className="text-lg mb-2">
            <strong className="font-semibold">ID:</strong> {booking.data.id}
          </p>
          <p className="text-lg mb-2">
            <strong className="font-semibold">Start Date:</strong> {new Date(booking.data.startDate).toLocaleDateString()}
          </p>
          <p className="text-lg mb-2">
            <strong className="font-semibold">End Date:</strong> {new Date(booking.data.endDate).toLocaleDateString()}
          </p>
          <p className="text-lg mb-2">
            <strong className="font-semibold">Total Price:</strong> ${booking.data.totalPrice.toFixed(2)}
          </p>
          <p className="text-lg mb-2">
            <strong className="font-semibold">Description:</strong> {booking.data.description}
          </p>
          <p className="text-lg mb-2">
            <strong className="font-semibold">Location:</strong> {booking.data.location}
          </p>
          <p className="text-lg mb-2">
            <strong className="font-semibold">Need:</strong> {booking.data.need}
          </p>
          <p className="text-lg">
            <strong className="font-semibold">Phone:</strong> {booking.data.phone}
          </p>
        </div>
      </div>
    </>
  );
};

export default BookingDetailPage;
