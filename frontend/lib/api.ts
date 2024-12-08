export async function fetchCars() {
    const response = await fetch('http://localhost:8080/api/car'); // Replace with your actual API endpoint
    const result = await response.json();

    if (!result.status) {
        throw new Error('Failed to fetch cars');
    }

    return Array.isArray(result.data) ? result.data : [];
}
