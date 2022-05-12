import { render, screen, fireEvent} from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import App from '../App';
import '@testing-library/jest-dom/extend-expect';

describe("Change Language", ()=>{
    it('should change the language to English when Dropdown changed to english', () =>{
        render(<App/>);
        const dropwdown = screen.getByTestId('dropdown');
        const headline = screen.getByTestId('headline');
        userEvent.selectOptions(dropwdown, 'english');
        expect(headline).toHaveTextContent("Get insight that help your business grow");
   })
   it('should change the language to Spanish when Dropdown changed to spanish', () =>{
    render(<App/>);
    const dropwdown = screen.getByTestId('dropdown');
    const headline = screen.getByTestId('headline');
    userEvent.selectOptions(dropwdown, 'spanish');
    expect(headline).toHaveTextContent("Obtenga información que ayude a su negocio a crecer");
    })
    
   it('should change the language to Indonesian when Dropdown changed to indonesian', () =>{
    render(<App/>);
    const dropwdown = screen.getByTestId('dropdown');
    const headline = screen.getByTestId('headline');
    userEvent.selectOptions(dropwdown, 'indonesian');
    expect(headline).toHaveTextContent("Dapatkan wawasan yang membantu bisnismu berkembang");
    })
    

})