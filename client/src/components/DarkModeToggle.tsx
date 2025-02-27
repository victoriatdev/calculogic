import React from "react";
import { IoSunny, IoMoon } from "react-icons/io5";

const DarkModeToggle = () => {
    const [dark, setDark] = React.useState(false);

    return (
        <button className="hover:text-(--color-action) cursor-pointer" type="button" role="switch" onClick={() => {
            setDark(!dark);
            document.body.classList.toggle("dark");
        }}>
            {
                dark && <IoSunny />
            } 
            {
                !dark && <IoMoon />
            }
        </button>
    );
};

export default DarkModeToggle;