import React from 'react';

const Loading: React.FC = () => (
  <div className="centered-container">
    <img
      src={`${process.env.PUBLIC_URL}/loading.svg`}
      alt="Loading animation"
    />
  </div>
);

export default Loading;
