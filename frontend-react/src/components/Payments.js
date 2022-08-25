import React from "react"

const Payments = (props) => {
    const cartItems = props.cartItems
    const totalPrice = props.totalPrice
    return (
        <div className="cart-items">
          
          {cartItems.length === 0 && (
              <div className="cart-items-empty"> No items added</div>
          )}
  
          {cartItems.map((cartItem) => (
                  <div className="card">
                      <div>{cartItem.product.name}</div>
                      <div>Quantity: {cartItem.quantity}</div>
                  </div>
              ))}
  
          {cartItems.length >= 1 && (<div>
            <h2> Total price: {totalPrice} zl</h2> 
            <button className="save-payments-button" onClick={() => props.makePay()}>Pay</button>
            </div>
          )}
        </div>
        
      )
}

export default Payments;
  