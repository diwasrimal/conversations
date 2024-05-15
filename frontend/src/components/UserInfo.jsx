import "./UserInfo.css";
import UserIcon from "../assets/user.svg";

export default function UserInfo({ user }) {
    return (
        <div className="user-info-container">
            <div className="picture-holder">
                <img src={UserIcon} alt="User Icon" />
            </div>
            <p className="normal-text">{`${user.fullname}`}</p>
        </div>
    );
}
