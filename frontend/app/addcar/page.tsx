"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

interface FormData {
  name: string;
  brand: string;
  seat: number;
  transmission: string;
  fuel: string;
  luggage: boolean;
  insurance: boolean;
  year: number;
  pricePerDay: number;
  availability: boolean;
  categoryId: number;
  image?: File | null;
}

interface AddCarResponse {
  status: boolean;
  message: string;
  errors: null | string;
}

const AddCarPage: React.FC = () => {
  const [formData, setFormData] = useState<FormData>({
    name: "",
    brand: "",
    seat: 0,
    transmission: "",
    fuel: "",
    luggage: false,
    insurance: false,
    year: new Date().getFullYear(),
    pricePerDay: 0,
    availability: true,
    categoryId: 1, // Default category ID; adjust as needed
    image: null,
  });
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);
  const router = useRouter();

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleCheckboxChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, checked } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: checked,
    }));
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files ? e.target.files[0] : null;
    
    if (file) {
      // Validasi tipe file
      if (file.type !== "image/jpeg") {
        setError("Only .png files are allowed.");
        setFormData((prev) => ({
          ...prev,
          image: null,
        }));
        return;
      }
  
      // Validasi ukuran file (maksimal 2 MB)
      if (file.size > 2 * 1024 * 1024) {
        setError("File size must be less than 2 MB.");
        setFormData((prev) => ({
          ...prev,
          image: null,
        }));
        return;
      }
  
      setFormData((prev) => ({
        ...prev,
        image: file,
      }));
      setError(null); // Reset error jika validasi berhasil
    }
  };
  

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setSuccess(null);
  
    const formDataObj = new FormData();
    formDataObj.append("name", formData.name);
    formDataObj.append("brand", formData.brand);
    formDataObj.append("seat", formData.seat.toString());
    formDataObj.append("transmission", formData.transmission);
    formDataObj.append("fuel", formData.fuel);
    formDataObj.append("luggage", formData.luggage.toString());
    formDataObj.append("insurance", formData.insurance.toString());
    formDataObj.append("year", formData.year.toString());
    formDataObj.append("pricePerDay", formData.pricePerDay.toString());
    formDataObj.append("availability", formData.availability.toString());
    formDataObj.append("categoryId", formData.categoryId.toString());
  
    if (formData.image) {
      formDataObj.append("image", formData.image);
    }
  
    console.log("Payload dikirim:", Object.fromEntries(formDataObj.entries()));
  
    try {
      const token = localStorage.getItem("token");
      const response = await fetch("http://localhost:8080/api/car", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`,
        },
        body: formDataObj,
      });
  
      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || "Adding car failed");
      }
  
      const result: AddCarResponse = await response.json();
  
      if (!result.status) {
        throw new Error(result.message || "Adding car failed");
      }
  
      setSuccess("Car added successfully!");
      setFormData({
        name: "",
        brand: "",
        seat: 0,
        transmission: "",
        fuel: "",
        luggage: false,
        insurance: false,
        year: new Date().getFullYear(),
        pricePerDay: 0,
        availability: true,
        categoryId: 1,
        image: null,
      });
  
      setTimeout(() => router.push("/cars"), 3000);
    } catch (error) {
      setError((error as Error).message);
    }
  };
  

  return (
    <div style={{ maxWidth: "600px", margin: "50px auto" }}>
      <h1>Add Car</h1>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {success && <p style={{ color: "green" }}>{success}</p>}
      <form onSubmit={handleSubmit}>
        <div className="py-3">
          <label className="input input-bordered flex items-center gap-2">
            <input
              type="text"
              id="name"
              name="name" 
              className="grow"
              value={formData.name}
              placeholder="Car Name"
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <div className="py-3">
          <label className="input input-bordered flex items-center gap-2">
            <input
              type="text"
              id="brand"
              name="brand"
              className="grow"
              value={formData.brand}
              placeholder="Car Brand"
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <div className="py-3">
          <label className="input input-bordered flex items-center gap-2">
            <input
              type="number"
              id="seat"
              name="seat"
              className="grow"
              value={formData.seat}
              placeholder="Number of Seats"
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <div className="py-3">
          <label className="input input-bordered flex items-center gap-2">
            <input
              type="text"
              id="transmission"
              name="transmission"
              className="grow"
              value={formData.transmission}
              placeholder="Transmission"
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <div className="py-3">
          <label className="input input-bordered flex items-center gap-2">
            <input
              type="text"
              id="fuel"
              name="fuel"
              className="grow"
              value={formData.fuel}
              placeholder="Fuel Type"
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <div className="py-3">
          <label className="input input-bordered flex items-center gap-2">
            <input
              type="checkbox"
              id="luggage"
              name="luggage"
              checked={formData.luggage}
              onChange={handleCheckboxChange}
            />
            <span>Luggage</span>
          </label>
        </div>
        <div className="py-3">
          <label className="input input-bordered flex items-center gap-2">
            <input
              type="checkbox"
              id="insurance"
              name="insurance"
              checked={formData.insurance}
              onChange={handleCheckboxChange}
            />
            <span>Insurance</span>
          </label>
        </div>
        <div className="py-3">
          <label className="input input-bordered flex items-center gap-2">
            <input
              type="number"
              id="year"
              name="year"
              className="grow"
              value={formData.year}
              placeholder="Year"
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <div className="py-3">
          <label className="input input-bordered flex items-center gap-2">
            <input
              type="number"
              id="pricePerDay"
              name="pricePerDay"
              className="grow"
              value={formData.pricePerDay}
              placeholder="Price Per Day"
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <div className="py-3">
          <label className="input input-bordered flex items-center gap-2">
            <input
              type="number"
              id="categoryId"
              name="categoryId"
              className="grow"
              value={formData.categoryId}
              placeholder="Category ID"
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <div className="py-3">
          <label className="input input-bordered flex items-center gap-2">
            <input
              type="file"
              id="image"
              name="image"
              className="grow"
              onChange={handleFileChange}
            />
          </label>
        </div>
        <button className="py-2 btn btn-block" type="submit">
          Add Car
        </button>
      </form>
    </div>
  );
};

export default AddCarPage;
