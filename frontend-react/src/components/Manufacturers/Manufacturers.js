import React, { useEffect, useState } from 'react'
import ManufacturerSingle from './ManufacturerSingle';
import "./Manufacturers.css"

const Manufacturers = () => {
    const [manufacturers, setManufacturers] = useState([])

    useEffect(() => {
        fetch("http://localhost:8000/manufacturers").then((res) => res.json()).then((manufacturers) => {
            const newManufacturers = manufacturers.map((manufacturer) => {
                return manufacturer
            })
            setManufacturers(newManufacturers)
    })
    }, [])

    return(
        <div className="manufacturers">
            {manufacturers.map((item) => (
                <ManufacturerSingle manufacturer={item}/>
            ))}
        </div>)
}

export default Manufacturers;