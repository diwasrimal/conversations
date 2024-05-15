import "./Button.css";

export default function Button({ children, style, ...rest }) {
    // Set in ./Button.css
    const vars = {
        "--width": style?.width || "130px",
        "--height": style?.height || "40px",
    };
    return (
        <button className="button" style={vars} {...rest}>
            {children}
        </button>
    );
}
