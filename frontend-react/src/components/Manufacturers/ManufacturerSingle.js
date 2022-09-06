import { Link } from "react-router-dom";

const ManufacturerSingle = ({manufacturer}) => {
    return (
        <div className="manufacturer-list">
            <Link to={"/manufacturers/" + manufacturer.ID} className="manufacturer-name">{manufacturer.name}</Link>
            <div className="manufacturer-country">Country: {manufacturer.origin_country}</div>
        </div>
    )
}

export default ManufacturerSingle;