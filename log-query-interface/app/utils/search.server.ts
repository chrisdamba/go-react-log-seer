import {Client} from '@elastic/elasticsearch'

const searchLogs = async (query) => {
  try {
    const esClient = new Client({node: 'http://localhost:9200'})
    const {
      fullTextSearch,
      level,
      message,
      resourceId,
      timestampStart,
      timestampEnd,
      traceId,
      spanId,
      commit,
      parentResourceId,
    } = query
    const must = []

    if (fullTextSearch) must.push({match: {_all: fullTextSearch}})
    if (level) must.push({match: {level}})
    if (message) must.push({match: {message}})
    if (resourceId) must.push({match: {resourceId}})
    if (timestampStart && timestampEnd) {
      must.push({
        range: {
          timestamp: {
            gte: timestampStart,
            lte: timestampEnd,
          },
        },
      })
    }
    if (traceId) must.push({match: {traceId}})
    if (spanId) must.push({match: {spanId}})
    if (commit) must.push({match: {commit}})
    if (parentResourceId)
      must.push({match: {'metadata.parentResourceId': parentResourceId}})

    const {body} = await esClient.search({
      index: 'logs',
      body: {
        query: {
          bool: {
            must,
          },
        },
      },
    })

    return body.hits.hits
  } catch (error) {
    console.log('Error searching logs: ', error)
  }
}

export default searchLogs
