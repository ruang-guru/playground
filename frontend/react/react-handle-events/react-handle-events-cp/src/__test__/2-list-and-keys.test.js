/**
 * @jest-environment jsdom
 */
 import React from 'react';
 import axios from "axios";
 import { render, screen } from "@testing-library/react";
 
 import PodcastLists from '../Screens/PodcastLists';
 
 const Modal = (props) => {
   const { children, onHide, show, ...otherProps } = props;
   return (
     <div {
       ...otherProps
      }>{children}</div>
   );
 };
 
 Modal.Body = ({ children }) => (
   <div>{children}</div>
 );
 
 Modal.Header = ({ children }) => (
   <div>
     {children}
   </div>
 );
 
 Modal.Title = ({ children }) => (
   <div>
     {children}
   </div>
 );
 
 Modal.Footer = ({ children }) => (
   <div>
     {children}
   </div>
 );
 
 const Button = (props) => {
   const { children, ...otherProps } = props;
   return (
     <button {...otherProps}>{children}</button>
   );
 }
 
 jest
   .mock('axios')
   .mock('react-bootstrap', () => ({
       Modal: Modal,
       Button: Button
   }));

describe('List and Keys', () => {
  const mockPodcastList = [{
    title: "title1",
    episode: 1,
    status: 1,
    genre: "genre1",
    duration: 1,
    publisher: "publisher1",
    summary: "summary1",
    imageUrl: "imageUrl",
    id: 1
  }, {
    title: "title2",
    episode: 2,
    status: 2,
    genre: "genre2",
    duration: 2,
    publisher: "publisher2",
    summary: "summary2",
    imageUrl: "imageUrl",
    id: 2
  }, {
    title: "title3",
    episode: 3,
    status: 3,
    genre: "genre3",
    duration: 3,
    publisher: "publisher3",
    summary: "summary3",
    imageUrl: "imageUrl",
    id: 3
  }];

  it('should render Podcast List Item depends on API response', async () => {
    axios.get.mockResolvedValue({
      status: 200,
      data: mockPodcastList
    });

    render(<PodcastLists />);

    const listItems = await screen.findAllByLabelText("podcast-list-item")
    expect(listItems).toHaveLength(3)
  });
});