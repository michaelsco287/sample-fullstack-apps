import { Button, Icon} from "antd";
import React from "react";

const BookInfo = (props: any) => {
    return(
        <div
        className="bookInfo"
        style={{ maxWidth: '100%',maxHeight: '100%' }}
      >
        <h1>{props.bookTitle}</h1>
        <img src={props.bookCover} style={{maxWidth: '100%',maxHeight: '100%'}}/>
        <h2>{props.bookAuthors}</h2>
        <Button> I own this book</Button>
        <Button>
          <Icon type="heart" theme="filled" />I want this book
        </Button>
      </div>
    )
}
export default BookInfo;
