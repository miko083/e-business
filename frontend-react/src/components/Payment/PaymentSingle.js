const PaymentSingle = ({cartItem}) => {
    return (
        <div className="card">
            <div className="console-name">{cartItem.product.name}</div>
            <div className="quantity">Quantity: {cartItem.quantity}</div>
        </div>
    )
}

export default PaymentSingle;