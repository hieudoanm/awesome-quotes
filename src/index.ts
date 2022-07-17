import axios from 'axios';
import fs from 'fs';
import get from 'lodash/get';

type Quote = {
  author: string;
  content: string;
};

const main = async () => {
  let quotes: Quote[] = [];
  for (let page = 1; page <= 14; page++) {
    console.info('page', page);
    const url = `https://api.quotable.io/quotes?page=${page}&limit=150`;
    const response = await axios(url);
    const results = get(response, 'data.results', []).map(
      ({ author, content }: { author: string; content: string }) => {
        return { author, content };
      }
    );
    quotes = quotes.concat(results);
  }

  const markdown = `# Quotes (${quotes.length})

${quotes
  .map((quote) => `- ${quote.author} - ${quote.content}`)
  .sort()
  .join('\n')}
`;
  await fs.writeFileSync('./README.md', markdown);
};

main().catch((error) => console.error(error));
