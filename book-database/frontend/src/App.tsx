import React, { useState } from "react";
import logo from "./logo.png";
import "./App.css";
import axios from "axios";
import { Button, Icon, Row, Col } from "antd";
import Search from "antd/lib/input/Search";

const App = () => {
  const getBook = async (value: string) => {
    const isbn: string = value;
    const result = await axios(`http://localhost:8080/book?isbn=${isbn}`);
    setBookCover(result.data["CoverURL"]);
    setBookTitle(result.data["Title"]);
    setBookAuthors(result.data["Authors"]);
  };
  const [bookCover, setBookCover] = useState("empty");
  const [bookAuthors, setBookAuthors] = useState("empty");
  const [bookTitle, setBookTitle] = useState("empty");

  return (
    <div className="App">
<Row
  type="flex"
  style={{ alignItems: "center" }}
  justify="center"
  gutter={[100,200]}
>
        <Col span={8}   style={{ alignItems: "center" }} >
          <img src={logo}  className="App-logo" alt="logo" />{" "}
        </Col>
      </Row>
      <Row gutter={[8, 248]}>
                <Col span={8}>
          <Search
            placeholder="Search for a book by its ISBN"
            onSearch={value => getBook(value)}
            enterButton
          />
          <div
            className="bookInfo"
            style={{ display: bookCover != "empty" ? "block" : "none" }}
          >
            <h2>{bookTitle}</h2>
            <img src={bookCover} />
            <h2>{bookAuthors}</h2>
            <Button> I own this book</Button>
            <Button>
              <Icon type="heart" theme="filled" />I want this book
            </Button>
          </div>
        </Col>
        <Col span={8} />
        <Col span={8} />
      </Row>

      <div className="div3">
        <h1>Owned Books</h1>
      </div>
      <div className="div4">
        <h1>Wanted Books</h1>
      </div>
    </div>
  );
};

export default App;
