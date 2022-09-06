import React from "react"
import "./Cart.css"
import {useNavigate} from 'react-router-dom'
import CartSingleProduct from "./CartSingleProduct"

const Cart = ({cartItems, handleAddProduct, handleRemoveProduct, handleCartClearance}) => {
  const navigate = useNavigate()
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
    
    return (
      <div className="cart-items">
        
        {cartItems.length === 0 && (
            <div className="cart-items-empty"> No items added</div>
        )}

        {cartItems.map((cartItem) => (
                <CartSingleProduct cartItem={cartItem} handleAddProduct={handleAddProduct} handleRemoveProduct={handleRemoveProduct}/>
            ))}

        {cartItems.length >= 1 && (<div>
            <button className="clear-cart-button" onClick={() => handleCartClearance()}>Clear cart</button>
            <br></br>
            <button className="save-payments-button" onClick={() =>{submitCart(); navigate("/payments")}}>Save and go to payments</button>
            </div>
        )}
      </div>
      
    )
}

export default Cart;
  