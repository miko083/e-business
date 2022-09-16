import React from "react"
import "./Cart.css"
import CartSingleProduct from "./CartSingleProduct"

const Cart = ({submitCart, cartItems, getCart, handleAddProduct, handleRemoveProduct, handleCartClearance}) => {
    return (
      <div className="cart-items">
        
        {cartItems.length === 0 && (<div>
            <div className="cart-items-empty"> No items added</div>
            <button className="get-saved-button" onClick={() => getCart()}>Get saved data</button>
            </div>
        )}

        {cartItems.map((cartItem) => (
                <CartSingleProduct cartItem={cartItem} handleAddProduct={handleAddProduct} handleRemoveProduct={handleRemoveProduct}/>
            ))}

        {cartItems.length >= 1 && (<div>
            <button className="clear-cart-button" onClick={() => handleCartClearance()}>Clear cart</button>
            <br></br>
            <button className="save-payments-button" onClick={() =>{submitCart()}}>Save and go to payments</button>
            </div>
        )}
      </div>
      
    )
}

export default Cart;
  