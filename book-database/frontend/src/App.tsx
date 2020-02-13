import React, { useState, useEffect } from "react";
import "./App.css";
import axios from "axios";
import BookInfo from "./BookInfo"
import {Icon,  Layout, Menu, Spin} from "antd";
import Search from "antd/lib/input/Search";
const { Header, Footer, Sider, Content } = Layout;


function Iconz() {
  return (
    <svg
    width="1em" height="1em" 
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 284.267 163"
          >
      <path
        fill="#D1A103"
        d="M6.5 8.15q1.86.6 1.86 2.57 0 1.32-.86 2.06-.85.73-2.33.73H1.08v-9.8h3.69q2.8 0 2.8 2.37 0 1.45-1.07 2.07zM2.87 5.31v2.25h1.64q.66 0 .98-.27.32-.28.32-.85 0-.6-.32-.87-.32-.26-.98-.26H2.87zm2.03 6.6q.78 0 1.2-.31.41-.32.41-1.07 0-.74-.43-1.06-.42-.31-1.18-.31H2.87v2.75H4.9zm7.85 1.74q-1.1 0-1.91-.49-.82-.49-1.25-1.32-.43-.82-.43-1.82 0-.99.43-1.83.43-.84 1.25-1.33.81-.49 1.91-.49 1.11 0 1.92.49.81.49 1.25 1.33.43.84.43 1.83 0 1-.43 1.82-.44.83-1.25 1.32-.81.49-1.92.49zm0-1.43q.86 0 1.34-.6.48-.6.48-1.6t-.48-1.61q-.48-.61-1.34-.61-.85 0-1.33.61-.49.61-.49 1.61t.49 1.6q.48.6 1.33.6zm7.98 1.43q-1.1 0-1.91-.49-.82-.49-1.25-1.32-.43-.82-.43-1.82 0-.99.43-1.83.43-.84 1.25-1.33.81-.49 1.91-.49 1.11 0 1.92.49.81.49 1.25 1.33.43.84.43 1.83 0 1-.43 1.82-.44.83-1.25 1.32-.81.49-1.92.49zm0-1.43q.86 0 1.34-.6.48-.6.48-1.6t-.48-1.61q-.48-.61-1.34-.61-.85 0-1.33.61-.49.61-.49 1.61t.49 1.6q.48.6 1.33.6zm4.79 1.29V3.01h1.74v6.37l2.17-2.87h1.96l-2.55 3.36 2.6 3.64h-1.96l-2.22-3.11v3.11h-1.74zm8.79.14q-.67 0-1.22-.13-.55-.12-1.14-.39l.15-1.48q.59.35 1.1.52.51.18 1.06.18 1.06 0 1.06-.7 0-.38-.28-.58-.28-.21-1.01-.49-1.03-.37-1.52-.86-.49-.48-.49-1.31 0-.91.7-1.47.7-.57 1.89-.57t2.17.55l-.16 1.48q-.53-.36-1-.56-.47-.2-.98-.2-.45 0-.7.19-.24.18-.24.51 0 .27.14.44.15.17.42.3.26.13.88.37 1.02.39 1.47.86.45.47.45 1.25 0 1-.72 1.54-.71.55-2.03.55zm3.77-.14v-1.9h1.82v1.9h-1.82zm3.28-8.06V3.71h1.73v1.74h-1.73zm0 8.06v-7h1.73v7h-1.73zm6.53.14q-1.1 0-1.91-.49-.82-.49-1.25-1.32-.43-.82-.43-1.82 0-.99.43-1.83.43-.84 1.25-1.33.81-.49 1.91-.49 1.11 0 1.92.49.81.49 1.25 1.33.43.84.43 1.83 0 1-.43 1.82-.44.83-1.25 1.32-.81.49-1.92.49zm0-1.43q.86 0 1.34-.6.48-.6.48-1.6t-.48-1.61q-.48-.61-1.34-.61-.85 0-1.33.61-.49.61-.49 1.61t.49 1.6q.48.6 1.33.6z"
        transform="matrix(5.6391 0 0 5.6391 -6.09 86.026)"
      ></path>
      <path
        fill="#D1A103"
        stroke="none"
        d="M92.818 23.216V72.1s-13.907-3.91-24.24-1.923c-9.255 1.779-16.066 6.8-17.215 7.687v.148l-.096-.075-.095.075v-.148c-1.148-.887-7.959-5.907-17.211-7.687C23.626 68.189 8.357 72.1 8.357 72.1V23.071l-2.858.429v52.662s16.64-4.282 27.959-2.104c11.131 2.141 17.856 8.581 17.856 8.581s6.725-6.44 17.855-8.581c11.321-2.179 27.278 2.104 27.278 2.104V23.727c.001 0-1.388-.235-3.629-.511zM33.121 65.229c9.051 1.739 16.046 6.978 16.046 6.978V24.113s-7.032-5.245-16.104-6.989c-9.183-1.766-21.179 1.723-21.179 1.723v48.094s12.034-3.482 21.237-1.712zm56.163 1.888V19.023s-10.996-3.49-20.175-1.723c-9.072 1.743-16.335 6.987-16.335 6.987v48.095s7.223-5.237 16.276-6.979c9.202-1.767 20.234 1.714 20.234 1.714z"
        clipRule="evenodd"
        transform="matrix(1.36335 0 0 1.36335 73.005 -22.04)"
      ></path>
    </svg>
  );}

const App = () => {

  const getOwnedBooks = () => {
    axios.get(`http://localhost:8080/ownedbooks`)
    .then(
      function(response) {
        if (response.data) setOwnedBooks( response.data)
      }
    )
  }
  const getWantedBooks = () => {
    axios.get(`http://localhost:8080/wantedbooks`)
    .then(
      function(response) {
        if (response.data) setWantedBooks( response.data)
      }
    )
  }
  const getBook = async (value: string) => {
    const isbn: string = value;
    setLoadingIndicator(true)
    const result = await axios(`http://localhost:8080/book?isbn=${isbn}`);
    setLoadingIndicator(false)
    setBookISBN13(result.data["ISBN13"])
    setBookCover(result.data["CoverURL"]);
    setBookTitle(result.data["Title"]);
    setBookAuthors(result.data["Authors"]);
    setPage("BookInfo")
  };

  const handleClick = (e: any) => {
    setPage(e.key)
  };
  const [ownedBooks, setOwnedBooks] = useState([{CoverURL: ""}])
  const [wantedBooks, setWantedBooks] = useState([{CoverURL: ""}])
  const [bookISBN13, setBookISBN13] = useState("empty");
  const [loadingIndicator, setLoadingIndicator] = useState(false)
  const [bookCover, setBookCover] = useState("empty");
  const [currentPage, setPage] = useState("OwnedBooks");
  const [bookAuthors, setBookAuthors] = useState("empty");
  const [bookTitle, setBookTitle] = useState("empty");
  useEffect( () => { 
    getOwnedBooks()
    getWantedBooks()
  },[currentPage])
  return (
    <div className="App">
      <Layout style={{minHeight: '100vh'}}>
        <Header style={{ background: '#0B173D', padding: 0, minHeight:'100px'}}>
        <Icon style={{ fontSize: '100px', }}component={Iconz} />
        </Header>
        <Layout>
        <Sider style={{background: '#0B173D', overflow: 'auto'}}>      
        <Menu
          mode="inline"
          defaultSelectedKeys={['1']}
          onClick={handleClick}
          style={{height: '100%', borderRight: 0 }}
        >
                  <Search
            placeholder="Book by ISBN"
            onSearch={value => getBook(value)}
            enterButton
          />
          <Menu.Item key="OwnedBooks">Owned Books</Menu.Item>
          <Menu.Item key="CurrentlyReading">Currently Reading</Menu.Item>
          <Menu.Item key="WantedBooks">Wanted Books</Menu.Item>
          <Menu.Item key="AdvancedSearch">Advanced Search</Menu.Item>

        </Menu>  
  
          </Sider>
        <Content style ={{background: '#fff'}}>  
        <Spin style={{ margin:"auto",
  left:"0",
  right:"0",
  top:"0",
  bottom:"0",
  position:"fixed"}} size="large" tip="Loading..." spinning={loadingIndicator}>

        {currentPage === "BookInfo" ? <BookInfo getWantedBooks={getWantedBooks} getOwnedBooks={getOwnedBooks} bookISBN13={bookISBN13} bookCover={bookCover} bookAuthors ={bookAuthors} bookTitle={bookTitle} ownedBooks={ownedBooks} wantedBooks={wantedBooks} /> : ""}   
        {currentPage === "CurrentlyReading" ? <h1>Books I am currently reading</h1> : ""}
        {currentPage === "OwnedBooks" ? <div><h1>Owned Books</h1> <p>{Object.keys(ownedBooks).map(key => 
    <img src={ownedBooks[parseInt(key)].CoverURL}></img>
)}</p> </div>: ""}
        {currentPage === "WantedBooks" ? <div><h1>Wanted Books</h1> <p>{Object.keys(wantedBooks).map(key => 
    <img src={wantedBooks[parseInt(key)].CoverURL}></img>
)}</p> </div>: ""}
        {currentPage === "AdvancedSearch" ? <h1>Advanced Search</h1> : ""}
        </Spin>
        </Content>
      </Layout>
        <Footer>Created by Nickson Thanda</Footer>
      </Layout>
    </div>
  );
};

export default App;
