import React from 'react';
import { FaSearch } from "react-icons/fa";

const SearchBar = ({ filterText, onFilter, placeholderText }) => (
  <div className="search-bar-container">
    <div className="search-input-wrapper">
      <span
        className="search-icon"
        style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', width: '18px', height: '18px' }}
      >
        <FaSearch />
      </span>
      <input
        type="text"
        placeholder={placeholderText || "Search..."}
        value={filterText}
        onChange={onFilter}
        className="search-input"
      />
    </div>
  </div>
);

export default SearchBar;