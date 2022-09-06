import "./Cart.css"

const CartSingleProduct = ({handleAddProduct, handleRemoveProduct, cartItem}) => {
    return (
        <div className="item-in-cart">
            <div className="console-name">{cartItem.product.name}</div>
            <button className="cart-add-button" onClick={() => handleAddProduct(cartItem.product)}>+</button>
            <button className="cart-remove-button" onClick={() => handleRemoveProduct(cartItem.product)}>-</button>
            <div className="quantity">Quantity: {cartItem.quantity}</div>
        </div>
    )
}

export default CartSingleProduct