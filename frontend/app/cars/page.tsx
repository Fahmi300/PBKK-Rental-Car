"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";

// Define the Car interface
interface Car {
  id: number;
  name: string;
  brand: string;
  year: number;
  seat: number;
  insurance: boolean;
  fuel: string;
  luggage: boolean;
  transmission: string;
  price_per_day: number;
  availability: boolean;
  category_id: number;
}

interface CarResponse {
  status: boolean;
  message: string;
  data: Car[];
}

const CarPage: React.FC = () => {
  const [cars, setCars] = useState<Car[]>([]);
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  const fetchCars = async () => {
    setError(null);
    try {
      const response = await fetch("http://localhost:8080/api/car", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      });

      if (!response.ok) {
        throw new Error("Failed to fetch car data");
      }

      const result: CarResponse = await response.json();

      if (!result.status) {
        throw new Error(result.message || "Failed to fetch cars");
      }

      setCars(result.data);
    } catch (error) {
      setError((error as Error).message);
    }
  };

  useEffect(() => {
    fetchCars();
  }, []);

  return (
    <div className="container mx-auto p-6">
      <h1 className="text-3xl font-bold text-center mb-6">Car List</h1>
      {error && <p className="text-red-500 text-center">{error}</p>}
      <div className="flex justify-center mb-6">
        <button
          className="btn btn-primary"
          onClick={() => router.push("/addcar")}
        >
          Add New Car
        </button>
      </div>

      {/* Carousel Section */}
      {cars.map((car) => (
        <div key={car.id} className="mb-10">
          <h2 className="text-2xl font-semibold mb-4">{car.name}</h2>
          <div className="carousel w-full rounded-lg shadow-md mb-4">
            {/* Carousel Images */}
            <div id={`slide1-${car.id}`} className="carousel-item relative w-full">
              <img
                src={`https://via.placeholder.com/800x400?text=${car.name}+Image+1`}
                className="w-full"
                alt={`${car.name} Slide 1`}
              />
              <div className="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
                <a
                  href={`#slide3-${car.id}`}
                  className="btn btn-circle btn-secondary"
                >
                  ❮
                </a>
                <a
                  href={`#slide2-${car.id}`}
                  className="btn btn-circle btn-secondary"
                >
                  ❯
                </a>
              </div>
            </div>
            <div id={`slide2-${car.id}`} className="carousel-item relative w-full">
              <img
                src={`https://via.placeholder.com/800x400?text=${car.name}+Image+2`}
                className="w-full"
                alt={`${car.name} Slide 2`}
              />
              <div className="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
                <a
                  href={`#slide1-${car.id}`}
                  className="btn btn-circle btn-secondary"
                >
                  ❮
                </a>
                <a
                  href={`#slide3-${car.id}`}
                  className="btn btn-circle btn-secondary"
                >
                  ❯
                </a>
              </div>
            </div>
            <div id={`slide3-${car.id}`} className="carousel-item relative w-full">
              <img
                src={`https://via.placeholder.com/800x400?text=${car.name}+Image+3`}
                className="w-full"
                alt={`${car.name} Slide 3`}
              />
              <div className="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
                <a
                  href={`#slide2-${car.id}`}
                  className="btn btn-circle btn-secondary"
                >
                  ❮
                </a>
                <a
                  href={`#slide1-${car.id}`}
                  className="btn btn-circle btn-secondary"
                >
                  ❯
                </a>
              </div>
            </div>
          </div>

          {/* Car Information */}
          <div className="p-4 bg-white rounded-lg shadow-md">
            <p>
              <strong>Brand:</strong> {car.brand}
            </p>
            <p>
              <strong>Year:</strong> {car.year}
            </p>
            <p>
              <strong>Seats:</strong> {car.seat}
            </p>
            <p>
              <strong>Fuel:</strong> {car.fuel}
            </p>
            <p>
              <strong>Transmission:</strong> {car.transmission}
            </p>
            <p>
              <strong>Price per Day:</strong> ${car.price_per_day}
            </p>
            <p>
              <strong>Availability:</strong>{" "}
              {car.availability ? "Available" : "Not Available"}
            </p>
          </div>
        </div>
      ))}

      {/* Car List Section */}
      <table className="table-auto w-full mt-10 border-collapse border border-gray-300">
        <thead>
          <tr>
            <th className="border px-4 py-2">ID</th>
            <th className="border px-4 py-2">Name</th>
            <th className="border px-4 py-2">Brand</th>
            <th className="border px-4 py-2">Year</th>
            <th className="border px-4 py-2">Seats</th>
            <th className="border px-4 py-2">Fuel</th>
            <th className="border px-4 py-2">Transmission</th>
            <th className="border px-4 py-2">Price/Day</th>
            <th className="border px-4 py-2">Availability</th>
          </tr>
        </thead>
        <tbody>
          {cars.length === 0 ? (
            <tr>
              <td colSpan={9} className="text-center px-4 py-2">
                No cars available.
              </td>
            </tr>
          ) : (
            cars.map((car) => (
              <tr key={car.id}>
                <td className="border px-4 py-2">{car.id}</td>
                <td className="border px-4 py-2">{car.name}</td>
                <td className="border px-4 py-2">{car.brand}</td>
                <td className="border px-4 py-2">{car.year}</td>
                <td className="border px-4 py-2">{car.seat}</td>
                <td className="border px-4 py-2">{car.fuel}</td>
                <td className="border px-4 py-2">{car.transmission}</td>
                <td className="border px-4 py-2">${car.price_per_day}</td>
                <td className="border px-4 py-2">
                  {car.availability ? "Available" : "Not Available"}
                </td>
              </tr>
            ))
          )}
        </tbody>
      </table>
    </div>
  );
};

export default CarPage;
