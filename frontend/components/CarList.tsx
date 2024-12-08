import React from 'react';
import CarItem from './CarItem';

function CarList({ cars }) {
    return (
        <ul>
            {cars.map(car => (
                <CarItem key={car.id} car={car} />
            ))}
        </ul>
    );
}

export default CarList;
