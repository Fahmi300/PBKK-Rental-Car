'use client';

import React, { useEffect, useState } from "react";
import { useAuthRedirect } from "@/lib/auth";

interface UserProfile {
  name: string;
  email: string;
  role: string;
}

const UserProfilePage: React.FC = () => {
  useAuthRedirect(); // Cek autentikasi

  const [user, setUser] = useState<UserProfile | null>(null);

  useEffect(() => {
    const fetchUserProfile = async () => {
      try {
        const token = localStorage.getItem("token");
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
    return <div>Loading profile...</div>; // Loader saat data belum tersedia
  }

  return (
    <div style={styles.container}>
      <h1 style={styles.title}>User Profile</h1>
      <div style={styles.card}>
        <p><strong>Name:</strong> {user.name}</p>
        <p><strong>Email:</strong> {user.email}</p>
        <p><strong>Role:</strong> {user.role}</p>
      </div>
    </div>
  );
};

const styles = {
  container: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    padding: "2rem",
    fontFamily: "Arial, sans-serif",
  },
  title: {
    fontSize: "2rem",
    marginBottom: "1.5rem",
  },
  card: {
    border: "1px solid #ddd",
    borderRadius: "8px",
    padding: "1rem 2rem",
    boxShadow: "0 2px 4px rgba(0,0,0,0.1)",
    width: "100%",
    maxWidth: "400px",
    backgroundColor: "#f9f9f9",
  },
};

export default UserProfilePage;