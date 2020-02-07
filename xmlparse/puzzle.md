### Suggested Programming Problem

In a language of your choosing given an xml string, validate whether or not the xml tags match correctly. There can be 0 to n attributes associated with tags, so keep that in mind. Don't spend more than a few hours on this, we're not looking for perfection, more so general structure and tactic. Be prepared to discuss your solution in the interview.

An example of valid XML would be:

    <CD><TITLE area="greatest-hits">Greatest Hits</TITLE><ARTIST>Dolly Parton</ARTIST><COUNTRY>USA</COUNTRY><COMPANY>RCA</COMPANY><PRICE>9.90</PRICE><YEAR>1982</YEAR></CD>

An example of invalid xml would be:

    <CD><TITLE area="greatest-hits">Greatest Hits</TITLE><ARTIST>Dolly Parton</ARTIST><COUNTRY>USA</COUNTRY><COMPANY>RCA<PRICE>9.90</PRICE><YEAR>1982</YEAR></CD>

Note that Company is missing a closing tag.

Good luck!
