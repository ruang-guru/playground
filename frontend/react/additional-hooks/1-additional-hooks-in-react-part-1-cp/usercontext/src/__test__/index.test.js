import { render, screen, fireEvent} from "@testing-library/react";
import App from '../App';
import '@testing-library/jest-dom/extend-expect';

describe("Toggle Component", ()=>{
    it('should show "Toggle Dark Theme" when application started', () =>{
         render(<App/>);
         const toggleButtonText = screen.getByTestId('toggle-button-text');
         expect(toggleButtonText).toHaveTextContent("Toggle Dark Theme");
    })
    it('should show the light background container when application started', () =>{
        render(<App/>);
        const container = screen.getByTestId('container');
        expect(container).toHaveClass('light-theme');
    })
    it('should show the light color Header when application started', () =>{
        render(<App/>);
        const header = screen.getByTestId('header');
        expect(header).toHaveClass('light-theme');
    })
    it('should show "Toggle Light Theme" when toggle button is clicked', () =>{
        render(<App/>);
        const toggleButton = screen.getByTestId('toggle-button');
        fireEvent.click(toggleButton);
        const toggleButtonText = screen.getByTestId('toggle-button-text');
        expect(toggleButtonText).toHaveTextContent("Toggle Light Theme");
    })

    it('should change container to Dark Background when toggle button is clicked', () =>{
        render(<App/>);
        const toggleButton = screen.getByTestId('toggle-button');
        fireEvent.click(toggleButton);
        const container = screen.getByTestId('container');
        expect(container).toHaveClass('dark-theme');
    })

    it('should change the Header text to Dark Background when toggle button is clicked', () =>{
        render(<App/>);
        const toggleButton = screen.getByTestId('toggle-button');
        fireEvent.click(toggleButton);
        const header = screen.getByTestId('header');
        expect(header).toHaveClass('dark-theme');
    })

})