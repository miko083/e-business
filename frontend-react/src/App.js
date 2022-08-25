import React, { useEffect, useState } from 'react'
import './App.css';
import Header from './components/Header';
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom'

import Products from './components/Products';
import Cart from './components/Cart'
import Payments from './components/Payments'

import useLocalStorage from'./hooks/UseLocalStorage'

const App = () => {

  const [productsState, setProductsState] = useState([])
  const [cartItems, setCartItems] = useLocalStorage('cart_items', [])
  const [totalPrice, setTotalPrice] = useLocalStorage('total_price', 0)

  const handleAddProduct = (product) => {
    
    const productExist = cartItems.find((item) => item.product.ID === product.ID)
    if (productExist) { 
      setCartItems(
        cartItems.map((item) =>
         item.product === product 
         ? {...productExist, quantity: productExist.quantity + 1}
         : item)
      )
    }
    else {
      setCartItems([...cartItems, {product, quantity: 1}])
    }
    setTotalPrice(totalPrice + product.price)
    localStorage.setItem('cartItems', JSON.stringify(cartItems))
    }

  const handleRemoveProduct = (product) => {
    const productExist = cartItems.find((item) => item.product.ID === product.ID)
    if (productExist.quantity === 1){
      setCartItems(cartItems.filter((item) => item.product.ID !== product.ID))
    } else {
      setCartItems(
        cartItems.map((item) => item.product.ID === product.ID ? {...productExist, quantity: productExist.quantity - 1}: item)
      )
    }
    setTotalPrice(totalPrice - product.price)
    // localStorage.setItem('cartItems', cartItems)
  }

  const handleCartClearance = () => {
    setCartItems([])
    setTotalPrice(0)
    // localStorage.setItem('cartItems', cartItems)
  }

  // Data from backend
  useEffect(() => {
    fetch("http://localhost:8000/consoles").then((res) => res.json()).then((products) => {
        const newProducts = products.map((product) => {
            return product
        })
        setProductsState(newProducts)
    })
  }, [])
  
  const submitCart = () => {

    const consolesWithQuantityToSend = []

    cartItems.map((item) => consolesWithQuantityToSend.push({"console_id": item.product.ID,"quantity": item.quantity}))
    const dataToSend = {
      user_id: 1,
      consoles_with_quantity: consolesWithQuantityToSend
    }
    
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(dataToSend)
    }

    fetch(`http://localhost:8000/carts`, requestOptions);
  
  }

  const makePay = () => {

    const dataToSend = {
      user_id: 1,
      money_to_pay: totalPrice
    }
    
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(dataToSend)
    }

    fetch(`http://localhost:8000/payments`, requestOptions).then((response) => {
      if(!response.ok) throw new Error(response.status);
        else alert("Payment done.")
    })
    setCartItems([])
    setTotalPrice(0)

  }
  
  return (
    <div>
      <Router>
        <Header/>
        <Routes>
            <Route path="/" element={<Products productsItems={productsState} handleAddProduct={handleAddProduct}/>}></Route>
            <Route path="/cart" element={<Cart cartItems={cartItems} handleAddProduct={handleAddProduct} handleRemoveProduct={handleRemoveProduct} handleCartClearance={handleCartClearance} submitCart={submitCart}/>}></Route>
            <Route path="/payments" element={<Payments cartItems={cartItems} totalPrice={totalPrice} makePay={makePay}/>}></Route>
        </Routes>
      </Router>
    </div>
    );
  
}

export default App;
