import './Header.css';

const Header = ({ headerText }) => {
    return (
        <div className="header">
            <h1>{headerText}</h1>
        </div>
    );
};

export default Header