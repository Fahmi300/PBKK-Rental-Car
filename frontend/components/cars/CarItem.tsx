import React from 'react';

function CarItem({ car }) {
    return (
        <li>
            <h2>{car.name}</h2>
            <p>Brand: {car.brand}</p>
            <p>Year: {car.year}</p>
            <p>Price per Day: ${car.price_per_day}</p>
        </li>
    );
}

export default CarItem;
