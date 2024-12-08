"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";

interface UserRegisterDTO {
  name: string;
  email: string;
  password: string;
}

interface RegisterResponse {
  status: boolean;
  message: string;
  errors: null | string;
}

const RegisterPage: React.FC = () => {
  const [formData, setFormData] = useState<UserRegisterDTO>({
    name: "",
    email: "",
    password: "",
  });
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);
  const router = useRouter();

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setSuccess(null);

    try {
      const response = await fetch("http://localhost:8080/api/user", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || "Registration failed");
      }

      const result: RegisterResponse = await response.json();

      if (!result.status) {
        throw new Error(result.message || "Registration failed");
      }

      setSuccess("Registration successful! Please log in.");
      setFormData({ name: "", email: "", password: "" });

      // Optionally redirect to login page
      setTimeout(() => router.push("/login"), 3000);
    } catch (error) {
      setError((error as Error).message);
    }
  };

  return (
    <div style={{ maxWidth: "400px", margin: "50px auto" }}>
      <h1>Register</h1>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {success && <p style={{ color: "green" }}>{success}</p>}
      <form onSubmit={handleSubmit}>
        <div className="py-3">
          <label className="input input-bordered flex items-center gap-2">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 16 16"
              fill="currentColor"
              className="h-4 w-4 opacity-70"
            >
              <path d="M3 14s-1 0-1-1 1-4 6-4 6 3 6 4-1 1-1 1H3Zm3-6a3 3 0 1 0 0-6 3 3 0 0 0 0 6Z" />
            </svg>
            <input
              type="text"
              id="name"
              name="name"
              className="grow"
              value={formData.name}
              placeholder="Name"
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <div className="py-3">
          <label className="input input-bordered flex items-center gap-2">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 16 16"
              fill="currentColor"
              className="h-4 w-4 opacity-70"
            >
              <path d="M2.5 3A1.5 1.5 0 0 0 1 4.5v.793c.026.009.051.02.076.032L7.674 8.51c.206.1.446.1.652 0l6.598-3.185A.755.755 0 0 1 15 5.293V4.5A1.5 1.5 0 0 0 13.5 3h-11Z" />
              <path d="M15 6.954 8.978 9.86a2.25 2.25 0 0 1-1.956 0L1 6.954V11.5A1.5 1.5 0 0 0 2.5 13h11a1.5 1.5 0 0 0 1.5-1.5V6.954Z" />
            </svg>
            <input
              type="email"
              id="email"
              name="email"
              className="grow"
              value={formData.email}
              placeholder="Email"
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <div className="py-3">
          <label className="input input-bordered flex items-center gap-2">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 16 16"
              fill="currentColor"
              className="h-4 w-4 opacity-70"
            >
              <path
                fillRule="evenodd"
                d="M14 6a4 4 0 0 1-4.899 3.899l-1.955 1.955a.5.5 0 0 1-.353.146H5v1.5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5v-2.293a.5.5 0 0 1 .146-.353l3.955-3.955A4 4 0 1 1 14 6Zm-4-2a.75.75 0 0 0 0 1.5.5.5 0 0 1 .5.5.75.75 0 0 0 1.5 0 2 2 0 0 0-2-2Z"
                clipRule="evenodd"
              />
            </svg>
            <input
              type="password"
              id="password"
              name="password"
              className="grow"
              value={formData.password}
              placeholder="Password"
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <button className="py-2 btn btn-block" type="submit">
          Register
        </button>
      </form>
    </div>
  );
};

export default RegisterPage;
