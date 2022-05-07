import { render, screen } from '@testing-library/react';
import App from '../App';
import axios from 'axios';
jest.mock("axios")
describe('First Render', () => {
  axios.get.mockResolvedValue({ status: 200, data: {} })
  it('should render the components', () => {
    render(<App />);
    const app = screen.getByLabelText('App');
    expect(app).toBeInTheDocument();
  });

  it('should render greetings or app title', () => {
    render(<App />);
    const greetings = screen.getByLabelText("App Title");
    expect(greetings).toBeInTheDocument();
  });
})
