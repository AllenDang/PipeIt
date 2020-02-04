# PipeIt

PipeIt is a text transformation, conversion, cleansing and extraction tool.

<img src="https://github.com/AllenDang/PipeIt/blob/master/screenshot/screenshot1.png" alt="PipeIt screen shot1"/>

# Features

- Split - split text to text array by given separator.
- RegexpSplit - split text to text array by given regexp expression.
- Fields - Fields splits the string s around each instance of one or more consecutive white space characters.
- Match - filter text array by regexp.
- Replace - replace each element of a text array.
- Surround - add prefix or suffix to each lement of a text array.
- Trim - Trim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.
- Join - join text array to single line of text by given separator.
- Line - output text array line by line.

And more pipes are comming...

(More important, tell me your case will help me to create more pipes which will actually useful.)

PipeIt also supports to read from Stdin, so you could pipe data using "cat file | PipeIt".

# Usage

## Extract image links from a html source

<img src="https://github.com/AllenDang/PipeIt/blob/master/screenshot/findimageurl.gif" alt="PipeIt demo to find image urls from html"/>

## Add single quotation mark to every words

<img src="https://github.com/AllenDang/PipeIt/blob/master/screenshot/addquotation.gif" alt="PipeIt demo to add single quotation"/>

## Replace the comma separated string to lines

<img src="https://github.com/AllenDang/PipeIt/blob/master/screenshot/commatolines.gif" alt="PipeIt demo to replace comma"/>

# The reason for creating it

First of all, to test the GUI framework created by me, [giu](https://github.com/AllenDang/giu), for a real project.

It turns out giu is really useful for this kind of application. It just costs me 6 hours to build it from ground.

And I have this idea for years, to create a text process pipeline, to ease my daily text processing pain.

Hope it could be useful to you as well. :)
