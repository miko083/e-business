import './App.css';
import Header from './components/Header/Header';
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom'

import Products from './components/Products/Products';
import Cart from './components/Cart/Cart'
import Payments from './components/Payment/Payments'
import Login from './components/Login/Login';

import useLocalStorage from'./hooks/UseLocalStorage'
import React, {useState} from 'react'
import Manufacturers from './components/Manufacturers/Manufacturers';

import {backEndLink, frontEndLink, headersForRequests} from './components/RequestSetup'

const App = () => {

  const [cartItems, setCartItems] = useLocalStorage('cart_items', [])
  const [totalPrice, setTotalPrice] = useLocalStorage('total_price', 0)

  const [userEmail, setUserEmail] = useLocalStorage('email',"")
  const [isLoggedIn, setLoggedIn] = useLocalStorage('is_logged_in', false)
  const [loginToken, setLoginToken] = useLocalStorage('login_token', '')

  const [clientSecretStripe, setClientSecretStripe] = useState("")

  const errorTokenMessage = "Token mismatch. Please re-login and try again."
  const errorTokenShippingCartMessage = "Token mismatch or shipping cart is not saved. Please re-login and try again."

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
  
  const preparePayment = () => {

    const dataToSend = {
      user_email: userEmail,
      login_token: loginToken,
      money_to_pay: totalPrice
    }
    
    const requestOptions = {
      method: 'POST',
      headers: headersForRequests,
      body: JSON.stringify(dataToSend)
    }
    fetch(backEndLink + '/preparePayments', requestOptions).then((response) => {
      if(!response.ok) alert(errorTokenShippingCartMessage);
      else return response.json();
    }).then(response => {
          setClientSecretStripe(response.stripe_token)
    })
  }

  const makePayment = () => {

    const dataToSend = {
      user_email: userEmail,
      login_token: loginToken,
    }
    
    const requestOptions = {
      method: 'POST',
      headers: headersForRequests,
      body: JSON.stringify(dataToSend)
    }

    fetch(backEndLink + '/payments', requestOptions).then((response) => {
      if(!response.ok) alert(errorTokenShippingCartMessage);
      else return response.json();
    }).then(response => {
          alert("Payment done!")
          setCartItems([])
          setTotalPrice(0)
    })
  }

  const submitCart = () => {

    const consolesWithQuantityToSend = []

    cartItems.map((item) => consolesWithQuantityToSend.push({"console_id": item.product.ID,"quantity": item.quantity}))
    const dataToSend = {
      user_email: userEmail,
      login_token: loginToken,
      consoles_with_quantity: consolesWithQuantityToSend
    }
    
    const requestOptions = {
      method: 'POST',
      headers: headersForRequests,
      body: JSON.stringify(dataToSend)
    }
    fetch(backEndLink + '/carts', requestOptions).then((response) => {
        if(!response.ok) alert(errorTokenMessage)
        else window.location.href = frontEndLink + "/payments"
    })
  
}

const getCart = () => {
  const dataToSend = {
    user_email: userEmail,
    login_token: loginToken,
  }
  
  const requestOptions = {
    method: 'POST',
    headers: headersForRequests,
    body: JSON.stringify(dataToSend)
  }
  let cartItemsFromBackend = []
  let tempTotalPrice = 0
  setTotalPrice(0)

  fetch(backEndLink + '/cartsUser', requestOptions).then((response) => {
    if(!response.ok) alert(errorTokenMessage);
    else return response.json();
  }).then(response => {
    response["consoles_with_quantity"].map(consoleWithQuantity=> {
      cartItemsFromBackend.push({"product": consoleWithQuantity.console, "quantity": consoleWithQuantity.quantity }) ;
      tempTotalPrice += (consoleWithQuantity.console.price * consoleWithQuantity.quantity);
      return ""}); 
    
    setTotalPrice(tempTotalPrice)
    setCartItems(cartItemsFromBackend)
  })
}

  const logout = () => {
    const dataToSend = {
      user_email: userEmail,
      login_token: loginToken
    }

    const requestOptions = {
      method: 'POST',
      headers: headersForRequests,
      body: JSON.stringify(dataToSend)
    }

    fetch(backEndLink + '/logout', requestOptions).then((response) => {
      if(!response.ok) alert("Something went wrong! Please re-login!")
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
            <Route path="/cart" element={<Cart cartItems={cartItems} submitCart={submitCart} getCart={getCart} handleAddProduct={handleAddProduct} handleRemoveProduct={handleRemoveProduct} handleCartClearance={handleCartClearance}/>}></Route>
            <Route path="/payments" element={<Payments cartItems={cartItems} totalPrice={totalPrice} makePayment={makePayment} preparePayment={preparePayment} clientSecret={clientSecretStripe}/>}></Route>
        </Routes>
      </Router>
    </div>
    );
  
}

export default App;
