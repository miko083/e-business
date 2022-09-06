import './App.css';
import Header from './components/Header/Header';
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom'

import Products from './components/Products/Products';
import Cart from './components/Cart/Cart'
import Payments from './components/Payment/Payments'
import Login from './components/Login/Login';

import useLocalStorage from'./hooks/UseLocalStorage'
import React from 'react'
import Manufacturers from './components/Manufacturers/Manufacturers';

const App = () => {

  const [cartItems, setCartItems] = useLocalStorage('cart_items', [])
  const [totalPrice, setTotalPrice] = useLocalStorage('total_price', 0)

  const [userEmail, setUserEmail] = useLocalStorage('email',"")
  const [isLoggedIn, setLoggedIn] = useLocalStorage('is_logged_in', false)
  const [loginToken, setLoginToken] = useLocalStorage('login_token', '')
  
  const handleAddProduct = (product) => {
    const productExist = cartItems.find((item) => item.product.ID === product.ID)
    if (productExist) {
      productExist["quantity"] += 1
      setCartItems(cartItems)
    }
    else {
      setCartItems([...cartItems, {product, quantity: 1}])
    }
    setTotalPrice(totalPrice + product.price)
    }
  
  const handleCartClearance = () => {
      setCartItems([])
      setTotalPrice(0)
  
  }

  const handleRemoveProduct = (product) => {
      const productExist = cartItems.find((item) => item.product.ID === product.ID)
      if (productExist.quantity === 1){
        setCartItems(cartItems.filter((item) => item.product.ID !== product.ID))
      } else {
        productExist["quantity"] -= 1
        setCartItems(cartItems)
      }
      setTotalPrice(totalPrice - product.price)
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

  const logout = () => {
    const dataToSend = {
      email: userEmail,
      login_token: loginToken
    }

    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(dataToSend)
    }

    fetch(`http://localhost:8000/logout`, requestOptions).then((response) => {
      if(!response.ok) alert("Something went wrong!")
        else alert(userEmail + " log out.")
    })

    setCartItems([])
    setTotalPrice(0)
    setLoggedIn(false)
    setUserEmail("")
    setLoginToken("")
  }
  
  return (
    <div>
      <Router>
        <Header loginEmail={userEmail} isLoggedIn={isLoggedIn} logout={logout}/>
        <Routes>
            <Route path="/" element={<Products handleAddProduct={handleAddProduct} setUserEmail={setUserEmail} setLoggedIn={setLoggedIn} setLoginToken={setLoginToken} isLoggedIn={isLoggedIn}/>}></Route>
            <Route path="/manufacturers" element={<Manufacturers/>}></Route>
            <Route path="/manufacturers/:id" element={<Products handleAddProduct={handleAddProduct} setUserEmail={setUserEmail} setLoggedIn={setLoggedIn} setLoginToken={setLoginToken} isLoggedIn={isLoggedIn}/>}></Route>
            <Route path="/login" element={<Login/>}></Route> 
            <Route path="/cart" element={<Cart cartItems={cartItems} handleAddProduct={handleAddProduct} handleRemoveProduct={handleRemoveProduct} handleCartClearance={handleCartClearance}/>}></Route>
            <Route path="/payments" element={<Payments cartItems={cartItems} totalPrice={totalPrice} makePay={makePay}/>}></Route>
        </Routes>
      </Router>
    </div>
    );
  
}

export default App;
