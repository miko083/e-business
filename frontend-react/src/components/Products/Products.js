import "./Products.css"
import {useSearchParams, useParams} from 'react-router-dom'
import jwt_decode from "jwt-decode";
import ProductSingle from "./ProductSingle";
import React, { useEffect, useState} from 'react'


const Products = ({handleAddProduct, isLoggedIn, setUserEmail, setLoggedIn, setLoginToken}) =>{
    const [searchParams, setSearchParams] = useSearchParams()
    const token = searchParams.get("login_token")
    const [productsItems, setProductsItems] = useState([])
    const parameter = useParams()

    useEffect(() => {
        fetch("http://localhost:8000/consoles").then((res) => res.json()).then((products) => {
            const newProducts = products.map((product) => {
                return product
            })
            setProductsItems(newProducts)
      })
      }, [])
    
    var itemsToPresent = productsItems
    if (parameter.id != null){
        itemsToPresent = productsItems.filter( (product) => product.manufacturer_id == parameter.id)
    }
    
    if (token != null) {
        const loginDetails= jwt_decode(token);
        setUserEmail(loginDetails.email)
        setLoggedIn(true)
        setLoginToken(token)
    }

    return(
        <div className="products">
            {itemsToPresent.map((productItem) => (
                <ProductSingle productItem={productItem} handleAddProduct={handleAddProduct} isLoggedIn={isLoggedIn}/>
            ))}
        </div>)
}

export default Products;