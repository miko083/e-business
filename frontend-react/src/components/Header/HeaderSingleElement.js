import {Link} from "react-router-dom"
const HeaderSingleElement = ({link, nameToDisplay, functionForClick}) => {
    if (functionForClick == null){
        return (
        <ul>
            <li>
                <Link to={link} className='cart-link'>{nameToDisplay}</Link>
            </li>
        </ul>
        )
    }
    return (
        <ul>
            <li>
                <Link to={link} className='cart-link' onClick={functionForClick}>{nameToDisplay}</Link>
            </li>
        </ul>
        )
}

export default HeaderSingleElement;