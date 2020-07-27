import React from 'react';
import './LandingPageLinks.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faGithub } from '@fortawesome/free-brands-svg-icons';

const LandingPageLinks: React.FC = () => (
  <div id="landing-page-links">
    <a
      href="https://github.com/nchaloult/codenames"
      target="_blank"
      rel="noopener noreferrer"
      title="View the project on GitHub"
      aria-label="View the project on GitHub">
      <FontAwesomeIcon
        className="external-link-icon"
        icon={faGithub}
        size="2x"
      />
    </a>
  </div>
);

export default LandingPageLinks;
