import * as React from "react";
import * as ReactDOM from "react-dom";
import Cards from "./Cards";

test('testing', () => {
    const root = document.createElement("div");
    ReactDOM.render(<Cards />, root);
    expect(root.querySelector("h1").textContent).toBe("Check Out the Top Suggestions! ");
})