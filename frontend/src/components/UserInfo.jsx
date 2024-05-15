import "./UserInfo.css";

export default function UserInfo({ user }) {
    return (
        <div className="user-info-container">
            <div className="picture-holder">
                <i className="fa-regular fa-user"></i>
            </div>
            <p className="normal-text">{`${user.fullname}`}</p>
        </div>
    );
}
