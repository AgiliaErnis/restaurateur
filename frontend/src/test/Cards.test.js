import * as React from "react";
import * as ReactDOM from "react-dom";
import Cards from "../components/cards/Cards";

test('renders h1', () => {
    const root = document.createElement("div");
    ReactDOM.render(<Cards />, root);
    expect(root.querySelector("h1").textContent).toBe("Check Out the Top Suggestions! ");
})