import "./Button.css";

export default function Button({ children, clickHandler, style, ...rest }) {
    // Set in ./Button.css
    const vars = {
        "--width": style?.width || "130px",
        "--height": style?.height || "40px",
    };
    return (
        <button
            className="button"
            style={vars}
            onClick={clickHandler}
            {...rest}
        >
            {children}
        </button>
    );
}
