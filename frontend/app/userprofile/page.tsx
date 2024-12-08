'use client';

import React, { useEffect, useState } from "react";
import { useAuthRedirect } from "@/lib/auth";
import Navbar from "@/components/navbar"; // Import Navbar component

interface UserProfile {
    data: {
        name: string;
        email: string;
        role: string;
    };
}

const UserProfilePage: React.FC = () => {
  useAuthRedirect(); // Cek autentikasi

  const [user, setUser] = useState<UserProfile | null>(null);

  useEffect(() => {
    const fetchUserProfile = async () => {
      try {
        const token = localStorage.getItem("token");
        console.log(token);
        const response = await fetch("http://localhost:8080/api/user/me", {
          headers: {
            Authorization: `Bearer ${token}`, // Kirim token di header
          },
        });

        if (!response.ok) {
          throw new Error("Failed to fetch user profile");
        }

        const data: UserProfile = await response.json();
        setUser(data);
      } catch (error) {
        console.error(error);
      }
    };

    fetchUserProfile();
  }, []);

  if (!user) {
    return <div className="text-center mt-10 text-lg text-gray-700">Loading profile...</div>; // Loader saat data belum tersedia
  }
  console.log(user)

  return (
    <>
    <Navbar />
    <div className="flex flex-col items-center p-8 font-sans">
      <h1 className="text-3xl font-bold mb-6">User Profile</h1>
      <div className="border border-gray-300 rounded-lg p-6 shadow-md w-full max-w-md bg-gray-50">
        <p className="text-lg mb-2">
          <strong className="font-semibold">Name:</strong> {user.data.name}
        </p>
        <p className="text-lg mb-2">
          <strong className="font-semibold">Email:</strong> {user.data.email}
        </p>
        <p className="text-lg">
          <strong className="font-semibold">Role:</strong> {user.data.role}
        </p>
      </div>
    </div>
    </>
  );
};

export default UserProfilePage;
