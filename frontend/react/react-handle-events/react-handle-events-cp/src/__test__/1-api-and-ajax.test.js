/**
 * @jest-environment jsdom
 */
import React from 'react';
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import axios from "axios";

import PodcastLists from '../Screens/PodcastLists';
import Constants from '../Constants';
import PodcastFormModal from '../components/PodcastFormModal';

const Modal = (props) => {
  const { children, onHide, show, ...otherProps } = props;
  return (
    <div {...otherProps}>{children}</div>
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
  .mock('react', () => ({
    ...jest.requireActual('react'),
    useState: jest.fn()
  }))
  .mock('axios')
  .mock('react-bootstrap', () => ({
      Modal: Modal,
      Button: Button
  }))
  .mock('../components/PodcastListItem', () => 'PodcastListItem');

describe('API and AJAX in Podcast List', () => {
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
  }];
  const mockPodcastFormModalProps = {
    podcastList: [],
    setPodcastList: jest.fn(),
    showFormModal: true,
    setShowFormModal: jest.fn(),
    setFormModalType: jest.fn(),
    formModalType: "ADD",
    podcastId: 1
  }
  const mockSetPodcastList = jest.fn();
  const mockSetShowFormModal = jest.fn();

  afterEach(() => {
    jest.clearAllMocks()
  })

  it('should call get podcast list and set to podcastList state', async () => {
    React.useState.mockImplementation(initValue => [initValue, mockSetPodcastList]);
    jest.spyOn(React, 'useEffect').mockImplementationOnce(f => f());

    axios.get.mockResolvedValue({
      status: 200,
      data: mockPodcastList
    });

    render(<PodcastLists />);
    expect(axios.get).toHaveBeenCalledTimes(1);
    expect(axios.get).toHaveBeenCalledWith(Constants.API_URL);

    await waitFor(async () => {
      expect(mockSetPodcastList).toHaveBeenCalledWith(mockPodcastList);
    });
  });

  it('should call POST submit API when user click submit button on ADD modal', async () => {
    React.useState.mockImplementation(initValue => [initValue, mockSetShowFormModal]);
    jest.spyOn(React, 'useEffect').mockImplementationOnce(f => f());

    axios.post.mockResolvedValue({
      status: 200,
      data: mockPodcastList
    });

    render(
      <PodcastFormModal {...mockPodcastFormModalProps}/>
    );
    const submitBtn = screen.getByLabelText('submit-btn')
    fireEvent.click(submitBtn)
    
    await waitFor(async () => {
      expect(axios.post).toHaveBeenCalledTimes(1);
    });
    expect(axios.post).toHaveBeenCalledWith(Constants.API_URL);
    expect(mockPodcastFormModalProps.setPodcastList).toBeCalled()
  });
  it('should call PUT submit API when user click submit button on ADD modal', async () => {
    const mockUpdateModalProps = {
      ...mockPodcastFormModalProps,
      formModalType: "UPDATE"
    }
    React.useState.mockImplementation(initValue => [initValue, mockSetShowFormModal]);
    jest.spyOn(React, 'useEffect').mockImplementationOnce(f => f());

    axios.put.mockResolvedValue({
      status: 200,
      data: mockPodcastList
    });

    render(
      <PodcastFormModal {...mockUpdateModalProps}/>
    );
    const submitBtn = screen.getByLabelText('submit-btn')
    fireEvent.click(submitBtn)
    
    await waitFor(async () => {
      expect(axios.put).toHaveBeenCalledTimes(1);
    });
    expect(axios.put).toHaveBeenCalledWith(
      Constants.API_URL + "/" + mockUpdateModalProps.podcastId, {}
    );
    expect(mockUpdateModalProps.setPodcastList).toBeCalled()
  });
});
