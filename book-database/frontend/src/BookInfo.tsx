import { Button, Icon, Alert} from "antd";
import React, {useState, useEffect} from "react";
import axios from "axios";

const BookInfo = (props: any) => {
    const ownClick = async () => {
      setOwnLoadingIndicator(true)
      var bodyFormData : FormData = new FormData();
      bodyFormData.append("isbn", props.bookISBN)
      axios({
        method: 'post',
        url: 'http://localhost:8080/ownedbooks',
        data: bodyFormData,
        })
     setAlertMessage("You have succesfully added this book to your list")
      setOwnLoadingIndicator(false)
      setDisabledOwnedButton(true)

    }

    const wantClick = async () => {
      setWantLoadingIndicator(true)
      var bodyFormData : FormData = new FormData();
      bodyFormData.append("isbn", props.bookISBN)
      axios({
        method: 'post',
        url: 'http://localhost:8080/wantedbooks',
        data: bodyFormData,
        })
      setAlertMessage("You have succesfully added this book to your list")
      setWantLoadingIndicator(false)
      setDisabledWantedButton(true)

    }
    useEffect( () => {
      if (props.OwnedBooks.some((e: { bookISBN: any; }) => {
        return e.bookISBN === props.bookISBN;
      }))
      {
        setDisabledOwnedButton(true)
      }
      if (props.WantedBooks.some((e: { bookISBN: any; }) => e.bookISBN === props.bookISBN))
      {
        setDisabledWantedButton(true)
      }
      },[])

    const [disabledWantedButton, setDisabledWantedButton] = useState(false)
    const [disabledOwnedButton, setDisabledOwnedButton] = useState(false)
    const [alertMessage, setAlertMessage] = useState("")
    const [ownLoadingIndicator, setOwnLoadingIndicator] = useState(false)
    const [wantLoadingIndicator, setWantLoadingIndicator] = useState(false)

    return(
        <div
        className="bookInfo"
        style={{ maxWidth: '100%',maxHeight: '100%' }}
      >
        <h1>{props.bookTitle}</h1>
        <img src={props.bookCover} style={{maxWidth: '100%',maxHeight: '100%'}}/>
        <h2>{props.bookAuthors}</h2>
        <Button loading={ownLoadingIndicator} onClick={ownClick} disabled={disabledOwnedButton}> I own this book</Button>
        <Button loading={wantLoadingIndicator} onClick={wantClick} disabled={disabledWantedButton}>
          <Icon type="heart" theme="filled" />I want this book
        </Button>
        {alertMessage && <Alert message={alertMessage} type="success" />}
      </div>
    )
}
export default BookInfo;
