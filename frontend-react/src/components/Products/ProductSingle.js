const ProductSingle = ({productItem, handleAddProduct, isLoggedIn}) => {
    return (
        <div className="card">
                <div className="product-name">{productItem.name}</div>
                <div className="product-manufacturer">Manufacturer: {productItem.manufacturer.name}</div>
                <div className="product-price">Price: {productItem.price} zl</div>
                {isLoggedIn === true && (
                    <button className="product-add-button" onClick={() => handleAddProduct(productItem)}>Add to Cart</button>
                )}
                {isLoggedIn === false && (
                    <button className="login-button">Please login to add to the cart</button>
                )}
        </div>
    )
}

export default ProductSingle;