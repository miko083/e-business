import React from 'react'
import {Link} from "react-router-dom"
import "./Header.css"

const Header = () => {
  return (
    <header className='header'>
        <div>
            <h1>
                <Link to="/" className="logo">
                    Shopping cart for E-Biznes
                </Link>
            </h1>
        </div>
    <div className='header-links'>
        <ul>
            <li>
                <Link to="/" className='home-link'>Home</Link>
            </li>
        </ul>
        <ul>
            <li>
                <Link to="/cart" className='cart-link'>Cart</Link>
            </li>
        </ul>
        <ul>
            <li>
            <Link to="/payments"  className='payments-link'>Payments</Link>
            </li>
        </ul>
    </div>
    </header>)
}

export default Header;
