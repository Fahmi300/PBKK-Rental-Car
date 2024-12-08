'use client'
import React, { useEffect, useState } from 'react';
import CarList from '../../components/cars/CarList';
import Loader from '../../components/Loader';
import { fetchCars } from '../../lib/api';

function CarsPage() {
    const [cars, setCars] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        async function loadCars() {
            try {
                const data = await fetchCars();
                setCars(data);
            } catch (err) {
                setError("stri");
            } finally {
                setLoading(false);
            }
        }

        loadCars();
    }, []);

    if (loading) return <Loader />;
    if (error) return <p>{error}</p>;

    return (
        <div>
            <h1>Cars</h1>
            <CarList cars={cars} />
        </div>
    );
}

export default CarsPage;
