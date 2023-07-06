import got from 'got';

const HOST = 'https://httpbin.org' //'https://httpbin.dmuth.org'

describe('Httpbin', () => {
  it('get', async ({ }) => {
    const response = await got.post(`${HOST}/anything`, {
      data: {
        hello: 'world'
      }
    });

    console.log(response);
  });
});
