import { NavLink, Outlet } from "react-router-dom";
import HomeIcon from "../assets/house.svg";
import LogoutIcon from "../assets/logout.svg";
import SearchIcon from "../assets/search.svg";
import UsersIcon from "../assets/users.svg";
import "./BaseLayout.css";

// Represents a base layout of the page including nav bar for a logged in user.
// Pages like home, search can extend upon this layout.
// TODO: manage overflows due to scroll
export default function BaseLayout() {
    return (
        <div className="base">
            <NavBar />
            <main className="base-main-content">
                <Outlet />
            </main>
        </div>
    );
}

function NavBar() {
    return (
        <aside>
            <nav>
                <ul>
                    <li>
                        <NavLink to="/" title="Home">
                            <img src={HomeIcon} alt="Home Icon" />
                        </NavLink>
                    </li>
                    <li>
                        <NavLink to="/search" title="Search">
                            <img src={SearchIcon} alt="Search Icon" />
                        </NavLink>
                    </li>
                    <li>
                        <NavLink to="/people" title="People">
                            <img src={UsersIcon} alt="People Icon" />
                        </NavLink>
                    </li>
                    {/* Logout at last */}
                    <li>
                        <NavLink to="/logout" title="Logout">
                            <img src={LogoutIcon} alt="Logout Icon" />
                        </NavLink>
                    </li>
                </ul>
            </nav>
        </aside>
    );
}
