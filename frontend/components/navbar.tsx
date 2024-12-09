import React, { useEffect, useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";

const Navbar = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [isAdmin, setIsAdmin] = useState(false);
  const router = useRouter();

  useEffect(() => {
    // Check if the token exists in localStorage
    const token = localStorage.getItem("token");
    if (token) {
      setIsLoggedIn(true);

      // Fetch user role (simplified example)
      const userRole = localStorage.getItem("role");
      if (userRole === "admin") {
        setIsAdmin(true);
      }
    }
  }, []);

  const handleLogout = () => {
    // Remove token and role from localStor age
    localStorage.removeItem("token");
    localStorage.removeItem("role");

    // Reset states
    setIsLoggedIn(false);
    setIsAdmin(false);

    // Redirect to the login page
    router.push("/login");
  };

  return (
    <nav className="bg-gray-800 p-4">
      <div className="max-w-7xl mx-auto flex items-center justify-between">
        <Link href="/" className="text-white text-2xl font-bold">
          Car Rental
        </Link>
        <div className="flex space-x-4">
          <Link href="/" className="text-white hover:bg-gray-700 px-3 py-2 rounded-md text-sm font-medium">
            Rent
          </Link>
          <Link href="/cars" className="text-white hover:bg-gray-700 px-3 py-2 rounded-md text-sm font-medium">
            Cars
          </Link>
          {isAdmin && (
            <Link href="/addcar" className="text-white hover:bg-gray-700 px-3 py-2 rounded-md text-sm font-medium">
              Add Car
            </Link>
          )}
          <Link href="/userprofile" className="text-white hover:bg-gray-700 px-3 py-2 rounded-md text-sm font-medium">
            Profile
          </Link>
          {isLoggedIn ? (
            <button
              onClick={handleLogout}
              className="text-white hover:bg-gray-700 px-3 py-2 rounded-md text-sm font-medium"
            >
              Logout
            </button>
          ) : (
            <Link href="/login" className="text-white hover:bg-gray-700 px-3 py-2 rounded-md text-sm font-medium">
              Login
            </Link>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
