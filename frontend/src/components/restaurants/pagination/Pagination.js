import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Pagination from '@material-ui/lab/Pagination';
import './Pagination.css'

const useStyles = makeStyles((theme) => ({
  root: {
    '& > *': {
        margin: '4rem',
    },
  },
}));

export default function RestaurantPagination() {
    const classes = useStyles();
    return (
        <div className={classes.root}>
            <Pagination count={10} shape="rounded" size="large" showFirstButton showLastButton/>
        </div>
    )
}
