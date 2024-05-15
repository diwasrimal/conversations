import HomeIcon from "../assets/house.svg";
import LogoutIcon from "../assets/logout.svg";
import SearchIcon from "../assets/search.svg";
import "./BaseWithNav.css";

// Represents a base layout of the page including nav bar for a logged in user.
// Pages like home, search can extend upon this layout.
export default function BaseWithNav({ children }) {
    return (
        <div className="base">
            <NavBar />
            <main className="base-main-content">{children}</main>
        </div>
    );
}

function NavBar() {
    return (
        <aside>
            <nav>
                <ul>
                    <li>
                        <a href="/" title="Home">
                            <img src={HomeIcon} alt="Home Icon" />
                        </a>
                    </li>
                    <li>
                        <a href="/search" title="Search">
                            <img src={SearchIcon} alt="Search Icon" />
                        </a>
                    </li>
                    <li>
                        <a href="/logout" title="Logout">
                            <img src={LogoutIcon} alt="Logout Icon" />
                        </a>
                    </li>
                </ul>
            </nav>
        </aside>
    );
}
