import React, { useEffect, useState } from 'react'
import ManufacturerSingle from './ManufacturerSingle';
import "./Manufacturers.css"
import {backEndLink, headersForRequests} from '../RequestSetup'

const Manufacturers = () => {
    const [manufacturers, setManufacturers] = useState([])

    useEffect(() => {

        const requestOptions = {
            method: 'GET',
            headers: headersForRequests,
        }

        fetch(backEndLink + "/manufacturers", requestOptions).then((res) => res.json()).then((manufacturers) => {
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