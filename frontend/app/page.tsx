import { link } from "fs";
import Image from "next/image";

export default function Home() {
  return (
    <main>
      <h1>Dashboard</h1>
      <a href="/cars">Car</a>
    </main>
  );
}
