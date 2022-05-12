import { render, screen, fireEvent } from "@testing-library/react";
import App from "../App";
import '@testing-library/jest-dom/extend-expect';

describe("multiply", () =>{
    it('should able to multiply input one with input two', () =>{
        render(<App/>);
        const inputOne = screen.getByTestId('input-one');
        const inputTwo = screen.getByTestId('input-two');
        const answer = screen.getByTestId('answer');
        fireEvent.change(inputOne, { target: { value: '4' } });
        fireEvent.change(inputTwo, { target: { value: '2' } });
        expect(answer).toHaveTextContent("8");
   })

});