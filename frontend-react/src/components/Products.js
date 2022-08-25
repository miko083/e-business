import React from "react";
import "./Products.css"

const Products = (props) =>{
    const productsItems = props.productsItems
    

    return(
        
        <div className="products">
            {productsItems.map((productItem) => (
                <div className="card">
                    <div>{productItem.name}</div>
                    <div>Manufacturer: {productItem.manufacturer.name}</div>
                    <div>Price: {productItem.price} zl</div>
                    <button className="product-add-button" onClick={() => props.handleAddProduct(productItem)}>Add to Cart</button>
                </div>
            ))}
        </div>)
}

export default Products;