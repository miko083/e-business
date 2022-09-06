import React from 'react'
import {Link} from "react-router-dom"
import "./Header.css"
import HeaderSingleElement from './HeaderSingleElement'

const Header = ({isLoggedIn, loginEmail, logout}) => {
    return (
        <header className='header'>
            <div>
                <h1>
                    <Link to="/" className="logo">
                        Consoles Shop E-Biznes
                    </Link>
                </h1>
            </div>
        <div className='header-links'>
            {isLoggedIn === true && (
            <div>
                <div>Logged as {loginEmail}</div>
                <HeaderSingleElement link={"/cart"} nameToDisplay="Cart"/>
                <HeaderSingleElement link={"/payments"} nameToDisplay="Payments"/>
                <HeaderSingleElement link={"/"} nameToDisplay="Logout" functionForClick={logout}/>
            </div>
            )}
            {isLoggedIn === false && (
                <div>
                <HeaderSingleElement link={"/login"} nameToDisplay="Login"/>
                </div>
            )}
            <HeaderSingleElement link={"/"} nameToDisplay="Consoles"/>
            <HeaderSingleElement link={"/manufacturers"} nameToDisplay="Manufacturers"/>
        </div>
    </header>)
}

export default Header;
