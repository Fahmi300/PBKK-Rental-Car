import React from "react";
import Link from "next/link";

const Navbar = () => {
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
          <Link href="/about" className="text-white hover:bg-gray-700 px-3 py-2 rounded-md text-sm font-medium">
            Profile
          </Link>
          {/* If you have login/logout functionality */}
          <Link href="/login" className="text-white hover:bg-gray-700 px-3 py-2 rounded-md text-sm font-medium">
            Login
          </Link>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
